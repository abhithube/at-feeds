import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import { PRELOAD_DELAY } from '@/lib/constants'
import { feedsQueryOptions } from '@/lib/query'
import { Collection } from '@/lib/types'
import { QueryClient, useQueryClient } from '@tanstack/react-query'
import { FolderClosed, FolderOpen } from 'lucide-react'
import { useRef } from 'react'
import { CollectionItemList } from './collection-item-list'

type CollectionsAccordionProps = {
  collections: Collection[]
}

export function CollectionsAccordion({
  collections,
}: CollectionsAccordionProps) {
  const queryClient = useQueryClient()
  const delayHandler = useRef<ReturnType<typeof setTimeout>>()

  function handleMouseEnter(collection: Collection) {
    clearTimeout(delayHandler.current)

    delayHandler.current = setTimeout(
      () => loadCollection(queryClient, collection),
      PRELOAD_DELAY,
    )
  }

  function handleMouseLeave() {
    clearTimeout(delayHandler.current)
  }

  return (
    <Accordion type="multiple">
      {collections.map((collection) => (
        <AccordionItem
          key={collection.id}
          value={`${collection.id}`}
          onMouseEnter={() => handleMouseEnter(collection)}
          onMouseLeave={handleMouseLeave}
        >
          <AccordionTrigger className="group/trigger rounded-md pl-4 hover:bg-muted/50">
            <div className="flex items-center">
              <FolderClosed className="mr-4 h-4 w-4 shrink-0 group-[&[data-state=open]]/trigger:hidden" />
              <FolderOpen className="mr-4 h-4 w-4 shrink-0 group-[&[data-state=closed]]/trigger:hidden" />
              <span className="text-sm">{collection.title}</span>
            </div>
          </AccordionTrigger>
          <AccordionContent>
            <CollectionItemList collection={collection} />
          </AccordionContent>
        </AccordionItem>
      ))}
    </Accordion>
  )
}

async function loadCollection(
  queryClient: QueryClient,
  collection: Collection,
) {
  await queryClient.ensureQueryData(
    feedsQueryOptions({
      collectionId: collection.id,
      limit: -1,
    }),
  )
}
