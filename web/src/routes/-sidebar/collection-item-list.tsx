import { collectionsQueryOptions, feedsQueryOptions } from '@/lib/query'
import { Collection } from '@/lib/types'
import { useQuery } from '@tanstack/react-query'
import { Loader2 } from 'lucide-react'
import { CollectionsAccordion } from './collections-accordion'
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

  const { data: collections } = useQuery(
    collectionsQueryOptions({
      parentId: collection.id,
      limit: -1,
    }),
  )

  if (!collections || !feeds) {
    return <Loader2 className="mx-auto mb-2 mt-1 animate-spin" />
  }

  return (
    <div className="space-y-1 pl-4">
      {collections.data.length > 0 && (
        <CollectionsAccordion collections={collections.data} />
      )}
      {feeds.data.map((feed) => (
        <FeedItem key={feed.id} feed={feed} />
      ))}
    </div>
  )
}
