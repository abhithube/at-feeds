import { LIMIT } from '@/lib/constants'
import { ensureInfiniteQueryData, feedEntriesQueryOptions } from '@/lib/query'
import { EntryGrid } from '@/routes/-entries/entry-grid'
import { useInfiniteQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'

export const Route = createFileRoute('/')({
  loader: async ({ context }) => {
    await ensureInfiniteQueryData(
      context.queryClient,
      feedEntriesQueryOptions({
        limit: LIMIT,
      }) as any,
    )
  },
  component: Component,
})

function Component() {
  const {
    data: entries,
    hasNextPage,
    fetchNextPage,
  } = useInfiniteQuery(
    feedEntriesQueryOptions({
      limit: LIMIT,
    }),
  )

  useEffect(() => {
    window.scrollTo(0, 0)
  }, [])

  if (!entries) return

  return (
    <div>
      <header className="fixed top-0 w-full bg-background p-8">
        <h1 className="text-3xl font-medium">All Feeds</h1>
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
