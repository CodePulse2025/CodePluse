import React from 'react';
import { Drawer, List, ListItem, ListItemText, Toolbar } from '@mui/material';

const menuItems = ['Home', 'Committers', 'PRs', 'Contributors'];

export default function Sidebar() {
  return (
    <Drawer variant="permanent" sx={{ width: 200 }}>
      <Toolbar />
      <List>
        {menuItems.map((text) => (
          <ListItem button key={text}>
            <ListItemText primary={text} />
          </ListItem>
        ))}
      </List>
    </Drawer>
  );
}
