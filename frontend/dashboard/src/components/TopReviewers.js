// src/components/TopReviewers.js
import React, { useEffect, useState } from 'react';
import { Card, CardContent, Typography, List, ListItem, ListItemText, CircularProgress, Box } from '@mui/material';
import axios from 'axios';

export default function TopReviewers() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    axios.get('http://localhost:8080/metrics/top-reviewers')
      .then((res) => {
        setData(res.data);
        setLoading(false);
      })
      .catch((err) => {
        console.error('Error fetching reviewers:', err);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return (
      <Card>
        <CardContent>
          <Typography variant="h6">Top Reviewers</Typography>
          <Box sx={{ display: 'flex', justifyContent: 'center', padding: 2 }}>
            <CircularProgress />
          </Box>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Top 3 Reviewers
        </Typography>
        <List>
          {data.map((reviewer, idx) => (
            <ListItem key={idx}>
              <ListItemText
                primary={`${reviewer.name || 'Unknown'}`}
                secondary={`Rank #${idx + 1} - ${reviewer.review_count} Reviews`}
              />
            </ListItem>
          ))}
        </List>
      </CardContent>
    </Card>
  );
}
