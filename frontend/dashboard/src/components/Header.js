import React from 'react';
import { AppBar, Toolbar, Typography, Box, InputBase, Select, MenuItem } from '@mui/material';

export default function Header() {
  return (
    <AppBar position="static" sx={{ backgroundColor: '#1976d2' }}>
      <Toolbar sx={{ display: 'flex', justifyContent: 'space-between' }}>
        <Typography variant="h6">CodePulse</Typography>
        
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Select size="small" defaultValue="repo1">
            <MenuItem value="repo1">Repo 1</MenuItem>
            <MenuItem value="repo2">Repo 2</MenuItem>
          </Select>
          <InputBase placeholder="Search…" sx={{ background: '#fff', px: 1, borderRadius: 1 }} />
        </Box>
      </Toolbar>
    </AppBar>
  );
}
