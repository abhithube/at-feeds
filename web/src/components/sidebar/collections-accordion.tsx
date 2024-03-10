import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import { Collection } from '@/lib/types'
import { FolderClosed, FolderOpen } from 'lucide-react'
import { CollectionItemList } from './collection-item-list'

type CollectionsAccordionProps = {
  collections: Collection[]
}

export function CollectionsAccordion({
  collections,
}: CollectionsAccordionProps) {
  return (
    <Accordion className="ml-4" type="multiple">
      {collections.map((collection) => (
        <AccordionItem key={collection.id} value={`${collection.id}`}>
          <AccordionTrigger className="group mr-2">
            <div className="flex items-center">
              <FolderClosed className="mr-4 h-4 w-4 group-[&[data-state=open]]:hidden" />
              <FolderOpen className="mr-4 h-4 w-4 group-[&[data-state=closed]]:hidden" />
              <span>{collection.title}</span>
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
