import { LIMIT } from '@/lib/constants'
import {
  ensureInfiniteQueryData,
  feedEntriesQueryOptions,
  feedQueryOptions,
} from '@/lib/query'
import { EntryGrid } from '@/routes/-entries/entry-grid'
import { useInfiniteQuery, useQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'

export const Route = createFileRoute('/feeds/$feedId')({
  loader: async ({ context, params }) => {
    await Promise.all([
      context.queryClient.ensureQueryData(feedQueryOptions(+params.feedId)),
      ensureInfiniteQueryData(
        context.queryClient,
        feedEntriesQueryOptions({
          feedId: +params.feedId,
          limit: LIMIT,
        }) as any,
      ),
    ])
  },
  component: Component,
})

function Component() {
  const { feedId } = Route.useParams()
  const { data: feed } = useQuery(feedQueryOptions(+feedId))
  const {
    data: entries,
    hasNextPage,
    fetchNextPage,
  } = useInfiniteQuery(
    feedEntriesQueryOptions({
      feedId: +feedId,
      limit: LIMIT,
    }),
  )

  useEffect(() => {
    window.scrollTo(0, 0)
  }, [feedId])

  if (!feed || !entries) return

  return (
    <div>
      <header className="fixed top-0 w-full bg-background p-8">
        <h1 className="text-3xl font-medium">{feed.title}</h1>
      </header>
      <main className="mt-24 h-full overflow-y-auto pb-8">
        <EntryGrid
          entries={entries.pages.map((page) => page.data).flat()}
          hasMore={hasNextPage}
          loadMore={fetchNextPage}
        />
      </main>
    </div>
  )
}
