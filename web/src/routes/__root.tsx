import { SettingsModal } from '@/components/settings/settings-modal'
import { AddFeedModal } from '@/components/sidebar/add-feed-modal'
import { Sidebar } from '@/components/sidebar/sidebar'
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/components/ui/resizable'
import { ensureInfiniteQueryData, feedsQueryOptions } from '@/lib/query'
import { QueryClient, useInfiniteQuery } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { Outlet, createRootRouteWithContext } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'

enum Modal {
  AddFeed = 'addFeed',
  Settings = 'settings',
}

export const Route = createRootRouteWithContext<{
  queryClient: QueryClient
}>()({
  loader: async ({ context }) => {
    await ensureInfiniteQueryData(
      context.queryClient,
      feedsQueryOptions({
        limit: -1,
      }) as any,
    )
  },
  component: Component,
})

function Component() {
  const { modal } = Route.useSearch<{ modal: Modal }>()

  const {
    data: feeds,
    hasNextPage,
    fetchNextPage,
  } = useInfiniteQuery(
    feedsQueryOptions({
      limit: -1,
    }),
  )

  if (!feeds) return

  return (
    <ResizablePanelGroup direction="horizontal">
      <ResizablePanel minSize={20} defaultSize={20} maxSize={30}>
        <div className="h-screen overflow-y-auto">
          <Sidebar
            feeds={feeds.pages[0].data}
            hasMore={hasNextPage}
            fetchMore={fetchNextPage}
          />
        </div>
      </ResizablePanel>
      <ResizableHandle />
      <ResizablePanel>
        <div className="h-screen overflow-y-auto">
          <Outlet />
        </div>
        {modal === Modal.AddFeed && <AddFeedModal />}
        {modal === Modal.Settings && <SettingsModal />}
        <TanStackRouterDevtools position="top-right" />
        <ReactQueryDevtools buttonPosition="bottom-right" />
      </ResizablePanel>
    </ResizablePanelGroup>
  )
}
