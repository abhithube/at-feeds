import { Separator } from '@/components/ui/separator'
import { Entry } from '@/lib/types'
import 'core-js/proposals/array-grouping-v2'
import { useInView } from 'react-intersection-observer'
import { EntryCard } from './entry-card'

type EntryGridProps = {
  entries: Entry[]
  hasMore: boolean
  loadMore?: () => void
}

export function EntryGrid({
  entries,
  hasMore = false,
  loadMore,
}: EntryGridProps) {
  const day = 1000 * 60 * 60 * 24
  const date = Date.now()
  const today = date - day
  const lastWeek = date - day * 7
  const lastMonth = date - day * 30
  const lastYear = date - day * 365

  const list = Object.entries(
    (Object as any).groupBy(entries, (item: Entry) => {
      const publishedAt = Date.parse(item.publishedAt)
      return publishedAt > today
        ? 'Today'
        : publishedAt > lastWeek
          ? 'This Week'
          : publishedAt > lastMonth
            ? 'This Month'
            : publishedAt > lastYear
              ? 'This Year'
              : 'This Lifetime'
    }) as { [key: string]: Entry[] },
  )

  const { ref } = useInView({
    threshold: 0,
    onChange: (inView) => {
      if (inView && loadMore) loadMore()
    },
  })

  return (
    <>
      <div className="space-y-8">
        {list.map(([title, entries], index) => (
          <div key={index} className="space-y-6">
            <div className="flex items-center space-x-8">
              <Separator className="flex-1" />
              <span className="text-sm font-medium text-muted-foreground">
                {title}
              </span>
              <Separator className="flex-1" />
            </div>
            <div className="space-y-6">
              {entries.map((entry, index) => (
                <div key={index} className="mx-auto max-w-[896px]">
                  <EntryCard entry={entry} />
                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
      {hasMore && <div ref={ref}></div>}
    </>
  )
}
