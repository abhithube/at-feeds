import { SettingsModal } from '@/components/settings/settings-modal'
import { AddFeedModal } from '@/components/sidebar/add-feed-modal'
import { Sidebar } from '@/components/sidebar/sidebar'
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/components/ui/resizable'
import { collectionsQueryOptions, feedsQueryOptions } from '@/lib/query'
import { QueryClient, useQuery } from '@tanstack/react-query'
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
    await Promise.all([
      context.queryClient.ensureQueryData(
        collectionsQueryOptions({
          parentId: -1,
          limit: -1,
        }),
      ),
      context.queryClient.ensureQueryData(
        feedsQueryOptions({
          collectionId: -1,
          limit: -1,
        }),
      ),
    ])
  },
  component: Component,
})

function Component() {
  const { modal } = Route.useSearch<{ modal: Modal }>()

  const { data: collections } = useQuery(
    collectionsQueryOptions({
      parentId: -1,
      limit: -1,
    }),
  )

  const { data: feeds } = useQuery(
    feedsQueryOptions({
      collectionId: -1,
      limit: -1,
    }),
  )

  if (!feeds || !collections) return

  return (
    <ResizablePanelGroup direction="horizontal">
      <ResizablePanel minSize={15} defaultSize={20} maxSize={30} collapsible>
        <div className="h-screen overflow-y-auto">
          <Sidebar collections={collections.data} feeds={feeds.data} />
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
