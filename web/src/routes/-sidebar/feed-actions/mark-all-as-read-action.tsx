import {
  DropdownMenuItem,
  DropdownMenuShortcut,
} from '@/components/ui/dropdown-menu'

export function MarkAllAsReadAction() {
  return (
    <DropdownMenuItem>
      Mark all as read
      <DropdownMenuShortcut>⇧⌘R</DropdownMenuShortcut>
    </DropdownMenuItem>
  )
}
