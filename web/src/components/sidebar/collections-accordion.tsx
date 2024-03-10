import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import { Collection } from '@/lib/types'
import { Folder } from 'lucide-react'
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
          <AccordionTrigger className="mr-2">
            <div className="flex items-center space-x-4">
              <Folder className="h-4 w-4" />
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
