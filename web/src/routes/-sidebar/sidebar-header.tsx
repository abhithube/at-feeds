import { Logo } from '@/components/logo'
import { Link } from '@tanstack/react-router'
import { Bookmark, Home } from 'lucide-react'

export function SidebarHeader() {
  return (
    <div>
      <div className="m-4">
        <Logo />
      </div>
      <div className="space-y-1 px-4">
        <Link
          className="group flex w-full items-center justify-between space-x-4 rounded-md px-4 py-2 text-sm font-medium text-primary hover:bg-muted"
          to="/"
          activeProps={{ className: 'bg-muted text-secondary' }}
        >
          <Home className="h-4 w-4 shrink-0" />
          <span className="grow truncate">Home</span>
        </Link>
        <Link
          className="group flex w-full items-center justify-between space-x-4 rounded-md px-4 py-2 text-sm font-medium text-primary hover:bg-muted"
          to="/saved"
          activeProps={{ className: 'bg-muted text-secondary' }}
        >
          <Bookmark className="h-4 w-4 shrink-0" />
          <span className="grow truncate">Saved</span>
        </Link>
      </div>
    </div>
  )
}
