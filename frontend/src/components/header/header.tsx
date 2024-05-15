"use client"

import * as React from "react"

import { cn } from "@/lib/utils"
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
  SmartWikiNavigationLink,
} from "@/components/ui/navigation-menu"

export default function Header() {
  return <>
    <NavigationMenu className="px-8">
      <span className="font-medium">SmartWiki</span>
      <NavigationMenuList>
        <NavigationMenuItem>
          <NavigationMenuTrigger>Getting started</NavigationMenuTrigger>
          <NavigationMenuContent>
            <ul className="grid gap-3 p-6 md:w-[400px] lg:w-[600px] lg:grid-cols-[.75fr_1fr]">
              <li className="row-span-3">
                <NavigationMenuLink asChild>
                  <a
                    className="flex h-full w-full select-none flex-col justify-end rounded-md bg-gradient-to-b from-muted/50 to-muted p-6 no-underline outline-none focus:shadow-md"
                    href="/"
                  >
                    <div className="mb-2 mt-4 text-lg font-medium">
                      SmartWiki
                    </div>
                    <p className="text-sm leading-tight text-muted-foreground">
                      A powerful and user-friendly wiki application featuring advanced search capabilities.
                    </p>
                  </a>
                </NavigationMenuLink>
              </li>
              <ListItem href="/docs/installation" title="Installation">
                Get started with SmartWiki in no time!
              </ListItem>
              <ListItem href="/docs/usage" title="Usage">
                Learn how to maximize your productivity with SmartWiki.
              </ListItem>
            </ul>
          </NavigationMenuContent>
        </NavigationMenuItem>
        <SmartWikiNavigationLink to="/demo" title="Demo" />
        <SmartWikiNavigationLink to="/docs" title="Documentation" />
      </NavigationMenuList>
    </NavigationMenu >
    <hr />
  </>
}



const ListItem = React.forwardRef<
  React.ElementRef<"a">,
  React.ComponentPropsWithoutRef<"a">
>(({ className, title, children, ...props }, ref) => {
  return (
    <li>
      <NavigationMenuLink asChild>
        <a
          ref={ref}
          className={cn(
            "block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
            className
          )}
          {...props}
        >
          <div className="text-sm font-medium leading-none">{title}</div>
          <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
            {children}
          </p>
        </a>
      </NavigationMenuLink>
    </li>
  )
})
ListItem.displayName = "ListItem"
