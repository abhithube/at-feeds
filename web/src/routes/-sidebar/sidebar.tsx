import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import { Collection, Feed } from '@/lib/types'
import { Link, useNavigate } from '@tanstack/react-router'
import { Bookmark, Home, Plus, Settings } from 'lucide-react'
import { Logo } from '../../components/logo'
import { CollectionsAccordion } from './collections-accordion'
import { FeedItem } from './feed-item'

type SidebarProps = {
  collections: Collection[]
  feeds: Feed[]
}

export const Sidebar = ({ collections, feeds }: SidebarProps) => {
  const navigate = useNavigate()

  return (
    <nav className="flex h-full flex-col">
      <div className="my-4 px-4">
        <Logo />
      </div>
      <div className="mb-4 space-y-1 px-4">
        <Link
          className="group flex w-full items-center justify-between space-x-4 rounded-md px-4 py-2 text-sm font-medium text-primary hover:bg-muted"
          to="/"
          activeProps={{ className: 'bg-muted text-secondary' }}
        >
          <Home className="h-4 w-4 shrink-0" />
          <span className="grow truncate">Home</span>
        </Link>
        <Link
          className="group flex w-full items-center justify-between space-x-4 rounded-md px-4 py-2 text-sm font-medium text-primary hover:bg-muted"
          to="/saved"
          activeProps={{ className: 'bg-muted text-secondary' }}
        >
          <Bookmark className="h-4 w-4 shrink-0" />
          <span className="grow truncate">Saved</span>
        </Link>
      </div>
      <Separator className="mb-2" />
      <div className="mb-2 flex w-full grow flex-col">
        <div className="flex items-center justify-between px-4">
          <span className="grow text-xs font-semibold text-muted-foreground">
            Feeds
          </span>
          <Button
            className="h-8 w-8 p-0"
            variant="ghost"
            onClick={() => {
              navigate({
                params: {},
                search: {
                  modal: 'addFeed',
                },
              })
            }}
          >
            <Plus className="h-4 w-4" />
          </Button>
        </div>
        <div className="space-y-1 px-4 pt-2">
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
      <div className="border-t p-4">
        <Button
          variant="ghost"
          className="group flex w-full items-center justify-between space-x-4 rounded-md px-4 py-2 text-sm font-medium text-primary"
          onClick={() => {
            navigate({
              params: {},
              search: {
                modal: 'settings',
              },
            })
          }}
        >
          <Settings className="h-4 w-4 shrink-0" />
          <span className="grow truncate text-left">Settings</span>
        </Button>
      </div>
    </nav>
  )
}