import {
  Inbox,
  LayoutDashboard,
  LucideIcon,
  Settings,
  Users,
} from 'lucide-react';

type Menu = {
  title: string;
  url: string;
};

type SidebarMenu = {
  icon: LucideIcon;
} & Menu;

export const sidebarMenus: SidebarMenu[] = [
  {
    title: 'Overview',
    icon: LayoutDashboard,
    url: '/dashboard/overview',
  },
  {
    title: 'People',
    icon: Users,
    url: '/dashboard/people',
  },
  {
    title: 'Inbox',
    icon: Inbox,
    url: '/dashboard/inbox',
  },
  {
    title: 'Settings',
    icon: Settings,
    url: '/dashboard/settings',
  },
];

export const navMenus: Menu[] = [
  {
    title: 'About',
    url: '/about',
  },
  {
    title: 'Features',
    url: '/features',
  },
  {
    title: 'Documentation',
    url: '/documentation',
  },
  {
    title: 'Contact',
    url: '/contact',
  },
];
