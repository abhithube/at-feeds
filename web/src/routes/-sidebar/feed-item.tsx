import { Feed } from '@/lib/types'
import { cn } from '@/lib/utils'
import { Link } from '@tanstack/react-router'
import { useState } from 'react'
import { Favicon } from '../../components/favicon'
import { FeedActions } from './feed-actions/feed-actions'

type FeedItemProps = {
  feed: Feed
}

export function FeedItem({ feed }: FeedItemProps) {
  const [open, setOpen] = useState(false)

  return (
    <Link
      key={feed.id}
      className="group flex w-full items-center justify-between space-x-4 rounded-md px-4 py-2 text-sm font-medium text-primary"
      activeProps={{
        className: 'bg-muted active text-secondary',
      }}
      inactiveProps={{
        className: 'hover:bg-muted/50',
      }}
      to="/feeds/$feedId"
      params={{
        feedId: `${feed.id}`,
      }}
    >
      <Favicon domain={new URL(feed.link).hostname} />
      <span className="grow truncate">{feed.title}</span>
      <div className="flex w-[3ch] shrink-0 justify-center">
        <span
          className={cn(
            'tabular-nums text-muted-foreground group-hover:hidden group-[.active]:text-secondary',
            feed.unreadEntryCount === 0 && 'hidden',
            open && 'hidden',
          )}
        >
          {feed.unreadEntryCount}
        </span>
        <FeedActions open={open} setOpen={setOpen} />
      </div>
    </Link>
  )
}
