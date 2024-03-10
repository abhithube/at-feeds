import { Rss } from 'lucide-react'

export function Logo() {
  return (
    <div className="flex items-center space-x-4">
      <Rss className="h-8 w-8 shrink-0 rounded-md bg-white bg-gradient-to-b from-secondary via-secondary to-secondary/60 p-1 text-white" />
      <span className="text-4xl font-bold">
        AT
        <span className="bg-white bg-gradient-to-b from-secondary via-secondary to-secondary/60 bg-clip-text text-transparent">
          Feeds
        </span>
      </span>
    </div>
  )
}
