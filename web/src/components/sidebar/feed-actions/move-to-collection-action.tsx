import {
  DropdownMenuItem,
  DropdownMenuPortal,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuSub,
  DropdownMenuSubContent,
  DropdownMenuSubTrigger,
} from '@/components/ui/dropdown-menu'

export function MoveToCollectionAction() {
  return (
    <DropdownMenuSub>
      <DropdownMenuSubTrigger>Move to collection</DropdownMenuSubTrigger>
      <DropdownMenuPortal>
        <DropdownMenuSubContent className="w-56">
          <DropdownMenuItem>Collection 1</DropdownMenuItem>
          <DropdownMenuItem>Collection 2</DropdownMenuItem>
          <DropdownMenuSeparator />
          <DropdownMenuItem>
            Create new
            <DropdownMenuShortcut>⇧⌘N</DropdownMenuShortcut>
          </DropdownMenuItem>
        </DropdownMenuSubContent>
      </DropdownMenuPortal>
    </DropdownMenuSub>
  )
}
