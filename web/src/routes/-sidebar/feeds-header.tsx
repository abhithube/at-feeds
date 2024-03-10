import { Button } from '@/components/ui/button'
import { useNavigate } from '@tanstack/react-router'
import { Plus } from 'lucide-react'

export function FeedsHeader() {
  const navigate = useNavigate()

  return (
    <div className="flex items-center justify-between px-4">
      <span className="grow text-xs font-semibold text-muted-foreground">
        Feeds
      </span>
      <Button
        className="h-8 w-8 p-0"
        variant="ghost"
        onClick={() => {
          navigate({
            search: {
              modal: 'addFeed',
            },
          })
        }}
      >
        <Plus className="h-4 w-4" />
      </Button>
    </div>
  )
}
