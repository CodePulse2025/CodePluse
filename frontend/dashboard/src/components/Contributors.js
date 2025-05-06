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

export default function Contributors() {
  const [contributors, setContributors] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    axios.get('http://localhost:8080/metrics/top-contributors')
      .then(res => {
        setContributors(res.data);
        setLoading(false);
      })
      .catch(err => {
        console.error('Error fetching contributors:', err);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return (
      <Card>
        <CardContent>
          <Typography variant="h6">Top 3 Contributors</Typography>
          <Box sx={{ display: 'flex', justifyContent: 'center', padding: 2 }}>
            <CircularProgress />
          </Box>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card sx={{ flex: 1 }}>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Top 3 Contributors
        </Typography>
        <List>
          {contributors.map((dev, idx) => (
            <ListItem key={idx}>
              <ListItemAvatar>
                <Avatar
                  alt={dev.Name}
                  src={`https://www.gravatar.com/avatar/${md5(dev.Email.trim().toLowerCase())}?d=identicon`}
                />
              </ListItemAvatar>
              <ListItemText
                primary={dev.Name || 'Unknown'}
                secondary={`Rank #${idx + 1}`}
              />
              <Chip
                label={`${dev.Count} PRs`}
                color="secondary"
                variant="outlined"
              />
            </ListItem>
          ))}
        </List>
      </CardContent>
    </Card>
  );
}
