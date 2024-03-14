import { feedsQueryOptions } from '@/lib/query'
import { Collection } from '@/lib/types'
import { useQuery } from '@tanstack/react-query'
import { Loader2 } from 'lucide-react'
import { FeedItem } from './feed-item'

type CollectionItemListProps = {
  collection: Collection
}

export function CollectionItemList({ collection }: CollectionItemListProps) {
  const { data: feeds } = useQuery(
    feedsQueryOptions({
      collectionId: collection.id,
      limit: -1,
    }),
  )

  if (!feeds) {
    return <Loader2 className="mx-auto mb-2 mt-1 animate-spin" />
  }

  return (
    <div className="relative space-y-1">
      <div className="ml-8 mt-1">
        {feeds.data.map((feed) => (
          <FeedItem key={feed.id} feed={feed} />
        ))}
      </div>
      <div className="absolute bottom-1 left-1.5 top-0 ml-4 border-l-[0.5px] border-muted-foreground/50"></div>
    </div>
  )
}
