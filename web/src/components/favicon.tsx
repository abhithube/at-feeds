export interface FaviconProps extends React.HTMLAttributes<HTMLDivElement> {
  domain: string
}

const Favicon = ({ className, domain, ...props }: FaviconProps) => {
  return (
    <img
      className={className}
      src={`https://icons.duckduckgo.com/ip3/${domain}.ico`}
      alt={domain}
      width={16}
      height={16}
      {...props}
    />
  )
}

export { Favicon }
