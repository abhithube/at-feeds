import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { cn } from '@/lib/utils'
import { MoreHorizontal } from 'lucide-react'
import { MarkAllAsReadAction } from './mark-all-as-read-action'
import { MoveToCollectionAction } from './move-to-collection-action'
import { UnsubscribeAction } from './unsubscribe-action'

type FeedActionsProps = {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export function FeedActions({ open, setOpen }: FeedActionsProps) {
  return (
    <DropdownMenu open={open} onOpenChange={setOpen}>
      <DropdownMenuTrigger>
        <MoreHorizontal
          className={cn(
            'hidden h-5 text-muted-foreground hover:text-primary group-hover:block',
            open && 'block text-primary',
          )}
        />
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56">
        <MoveToCollectionAction />
        <MarkAllAsReadAction />
        <DropdownMenuSeparator />
        <UnsubscribeAction />
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
