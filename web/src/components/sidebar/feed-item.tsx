import { Badge } from '@/components/ui/badge'
import { Feed } from '@/lib/types'
import { cn } from '@/lib/utils'
import { Link } from '@tanstack/react-router'
import { Favicon } from '../favicon'

type FeedItemProps = {
  feed: Feed
}

export function FeedItem({ feed }: FeedItemProps) {
  return (
    <Link
      key={feed.id}
      className="group flex w-full items-center justify-between space-x-4 rounded-md px-4 py-2 text-sm font-medium text-primary hover:bg-muted"
      activeProps={{ className: 'bg-muted text-secondary' }}
      to="/feeds/$feedId"
      params={{
        feedId: `${feed.id}`,
      }}
    >
      <Favicon domain={new URL(feed.link).hostname} />
      <span className="grow truncate">{feed.title}</span>
      <Badge
        variant="outline"
        className={cn(
          'py-[1px] text-secondary',
          feed.unreadCount === 0 && 'hidden',
        )}
      >
        {feed.unreadCount}
      </Badge>
    </Link>
  )
}
