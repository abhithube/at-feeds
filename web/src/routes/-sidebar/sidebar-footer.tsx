import { Button } from '@/components/ui/button'
import { useNavigate } from '@tanstack/react-router'
import { Settings } from 'lucide-react'

export function SidebarFooter() {
  const navigate = useNavigate()

  return (
    <div className="p-4 pt-0">
      <Button
        variant="ghost"
        className="flex w-full items-center justify-between space-x-4 rounded-md px-4 py-2 text-sm font-medium text-primary"
        onClick={() => {
          navigate({
            search: {
              modal: 'settings',
            },
          })
        }}
      >
        <Settings className="h-4 w-4 shrink-0" />
        <span className="grow truncate text-left">Settings</span>
      </Button>
    </div>
  )
}
