// src/components/TopCommitters.js
import React, { useEffect, useState } from 'react';
import {
  Card,
  CardContent,
  Typography,
  List,
  ListItem,
  ListItemAvatar,
  Avatar,
  ListItemText,
  Chip,
  CircularProgress,
  Box
} from '@mui/material';
import axios from 'axios';
import md5 from 'blueimp-md5';

export default function TopCommitters() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    axios.get('http://localhost:8080/metrics/top-committers')
      .then(res => {
        setData(res.data);
        setLoading(false);
      })
      .catch(err => {
        console.error('Error fetching committers:', err);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return (
      <Card>
        <CardContent>
          <Typography variant="h6">Top 3 Committers</Typography>
          <Box sx={{ display: 'flex', justifyContent: 'center', padding: 2 }}>
            <CircularProgress />
          </Box>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card sx={{ height: '100%' }}>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Top 3 Committers
        </Typography>
        <List>
          {data.map((dev, idx) => (
            <ListItem key={idx}>
                <ListItemAvatar>
                <Avatar
                  alt={dev.Name}
                  src={
                    dev.Name
                      ? `https://github.com/${dev.Name}.png`
                      : `https://www.gravatar.com/avatar/${md5(dev.Email?.trim().toLowerCase() || '')}?d=identicon`
                  }
                />

                </ListItemAvatar>
              <ListItemText
                primary={dev.Name || 'Unknown'}
                secondary={`Rank #${idx + 1}`}
              />
              <Chip
                label={`${dev.Count} commits`}
                color="primary"
                variant="outlined"
              />
            </ListItem>
          ))}
        </List>
      </CardContent>
    </Card>
  );
}
