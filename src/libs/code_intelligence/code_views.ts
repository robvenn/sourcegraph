import { from, merge, Observable, of, Subject } from 'rxjs'
import { filter, map, mergeMap } from 'rxjs/operators'

import { CodeHost, ResolvedCodeView } from './code_intelligence'

/**
 * Emits a ResolvedCodeView when it's DOM element is on or about to be on the page.
 */
const emitWhenIntersecting = (margin: number) => {
    const codeViewStash = new Map<HTMLElement, ResolvedCodeView>()

    const intersectingElements = new Subject<HTMLElement>()

    const intersectionObserver = new IntersectionObserver(
        entries => {
            for (const entry of entries) {
                // `entry` is an `IntersectionObserverEntry`,
                // which has
                // [isIntersecting](https://developer.mozilla.org/en-US/docs/Web/API/IntersectionObserverEntry/isIntersecting#Browser_compatibility)
                // as a prop, but TS complains that it does not
                // exist.
                if ((entry as any).isIntersecting) {
                    intersectingElements.next(entry.target as HTMLElement)
                }
            }
        },
        {
            rootMargin: `${margin}px`,
            threshold: 0,
        }
    )

    return (codeViews: Observable<ResolvedCodeView>) =>
        new Observable<ResolvedCodeView>(observer => {
            codeViews.subscribe(({ codeView, ...rest }) => {
                intersectionObserver.observe(codeView)
                codeViewStash.set(codeView, { codeView, ...rest })
            })

            intersectingElements
                .pipe(
                    map(element => codeViewStash.get(element)),
                    filter(codeView => !!codeView)
                )
                .subscribe(observer)
        })
}

/**
 * findCodeViews finds all the code views on a page given a CodeHost. It emits code views
 * that are lazily loaded as well.
 */
export const findCodeViews = (codeHost: CodeHost, watchChildrenModifications = true) => (
    containers: Observable<Element>
) => {
    const codeViewsFromList: Observable<ResolvedCodeView> = containers.pipe(
        filter(() => !!codeHost.codeViews),
        mergeMap(container =>
            from(codeHost.codeViews!).pipe(
                map(({ selector, ...info }) => ({
                    info,
                    matches: container.querySelectorAll<HTMLElement>(selector),
                }))
            )
        ),
        mergeMap(({ info, matches }) =>
            of(...matches).pipe(
                map(codeView => ({
                    ...info,
                    codeView,
                }))
            )
        )
    )

    const codeViewsFromResolver: Observable<ResolvedCodeView> = containers.pipe(
        filter(() => !!codeHost.codeViewResolver),
        map(container => ({
            resolveCodeView: codeHost.codeViewResolver!.resolveCodeView,
            matches: container.querySelectorAll<HTMLElement>(codeHost.codeViewResolver!.selector),
        })),
        mergeMap(({ resolveCodeView, matches }) =>
            of(...matches).pipe(
                map(codeView => ({
                    ...resolveCodeView(codeView),
                    codeView,
                }))
            )
        )
    )

    const obs = [codeViewsFromList, codeViewsFromResolver]

    if (watchChildrenModifications) {
        const possibleLazilyLoadedContainers = new Subject<HTMLElement>()

        const mutationObserver = new MutationObserver(mutations => {
            for (const mutation of mutations) {
                if (mutation.type === 'childList' && mutation.target instanceof HTMLElement) {
                    const { target } = mutation

                    possibleLazilyLoadedContainers.next(target)
                }
            }
        })

        containers.subscribe(container =>
            mutationObserver.observe(container, {
                childList: true,
                subtree: true,
            })
        )

        const lazilyLoadedCodeViews = possibleLazilyLoadedContainers.pipe(findCodeViews(codeHost, false))

        obs.push(lazilyLoadedCodeViews)
    }

    return merge(...obs).pipe(
        emitWhenIntersecting(250),
        filter(({ codeView }) => !codeView.classList.contains('sg-mounted'))
    )
}
