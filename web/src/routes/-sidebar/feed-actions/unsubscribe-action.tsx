import {
  DropdownMenuItem,
  DropdownMenuShortcut,
} from '@/components/ui/dropdown-menu'

export function UnsubscribeAction() {
  return (
    <DropdownMenuItem>
      Unsubscribe
      <DropdownMenuShortcut>⇧⌘U</DropdownMenuShortcut>
    </DropdownMenuItem>
  )
}
