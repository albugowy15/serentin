'use client';

import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from '@/components/ui/navigation-menu';
import { navMenus } from '@/config/sidebar';
import Link from 'next/link';

export default function HomeNavigation() {
  return (
    <NavigationMenu>
      <NavigationMenuList>
        {navMenus.map((menu, index) => (
          <NavigationMenuItem key={index}>
            <Link href={menu.url} legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                {menu.title}
              </NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
        ))}
      </NavigationMenuList>
    </NavigationMenu>
  );
}
