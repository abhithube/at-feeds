import {
  Card,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Separator } from '@/components/ui/separator'
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip'
import { formatRelativeDate } from '@/lib/format'
import { Entry } from '@/lib/types'
import { Favicon } from '../../components/favicon'
import { EntryCardActions } from './entry-card-actions'

type EntryCardProps = {
  entry: Entry
}

export function EntryCard({ entry }: EntryCardProps) {
  return (
    <Card className="overflow-hidden shadow-md">
      <div className="flex">
        <img
          className="aspect-video h-[165px] bg-background object-cover"
          src={entry.thumbnailUrl ?? 'https://placehold.co/320x180/black/black'}
          alt={entry.title}
          width={320}
          height={180}
          loading="lazy"
        />
        <div className="flex w-full flex-col">
          <CardHeader>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger>
                  <CardTitle className="line-clamp-1 text-left text-lg leading-tight">
                    <a href={entry.link} target="_blank">
                      {entry.title}
                    </a>
                  </CardTitle>
                </TooltipTrigger>
                <TooltipContent>{entry.title}</TooltipContent>
              </Tooltip>
            </TooltipProvider>
            <CardDescription className="line-clamp-2">
              {entry.content ?? 'No description available'}
            </CardDescription>
          </CardHeader>
          <CardFooter className="flex justify-between">
            <div className="flex h-4 space-x-2">
              <div className="flex space-x-1">
                <Favicon domain={new URL(entry.link).hostname} />
                <span className="text-xs font-semibold text-muted-foreground">
                  {entry.author ?? 'Anonymous'}
                </span>
              </div>
              <Separator
                className="bg-muted-foreground/50"
                orientation="vertical"
              />
              <span className="text-xs font-semibold text-muted-foreground">
                {formatRelativeDate(Date.parse(entry.publishedAt))}
              </span>
            </div>
            <EntryCardActions entry={entry} />
          </CardFooter>
        </div>
      </div>
    </Card>
  )
}
