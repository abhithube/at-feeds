import { Separator } from '@/components/ui/separator'
import { Collection, Feed } from '@/lib/types'
import { CollectionsAccordion } from './collections-accordion'
import { FeedItem } from './feed-item'
import { FeedsHeader } from './feeds-header'
import { SidebarFooter } from './sidebar-footer'
import { SidebarHeader } from './sidebar-header'

type SidebarProps = {
  collections: Collection[]
  feeds: Feed[]
}

export const Sidebar = ({ collections, feeds }: SidebarProps) => {
  return (
    <nav className="flex h-full flex-col space-y-4">
      <SidebarHeader />
      <Separator />
      <div className="flex w-full grow flex-col">
        <FeedsHeader />
        <div className="space-y-1 px-4">
          {feeds.length > 0 || collections.length > 0 ? (
            <>
              <CollectionsAccordion collections={collections} />
              {feeds.map((feed) => (
                <FeedItem key={feed.id} feed={feed} />
              ))}
            </>
          ) : (
            <div className="text-sm font-light">
              You have not subscribed to any feeds yet. Click + to add one.
            </div>
          )}
        </div>
      </div>
      <Separator />
      <SidebarFooter />
    </nav>
  )
}
