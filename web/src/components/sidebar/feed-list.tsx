import { Badge } from '@/components/ui/badge'
import { Feed } from '@/lib/types'
import { Link } from '@tanstack/react-router'
import { Favicon } from '../favicon'

type Props = {
  feeds: Feed[]
}

export function FeedList({ feeds }: Props) {
  return feeds.map((feed) => (
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
      {feed.unreadCount > 0 && (
        <Badge variant="outline" className="py-[1px] text-secondary">
          {feed.unreadCount}
        </Badge>
      )}
    </Link>
  ))
}
