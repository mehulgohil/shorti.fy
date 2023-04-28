import * as React from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import {ShortenURLS} from "../models/models";

interface BasicTableProps {
  allURLS: ShortenURLS[]
}

export default function BasicTable(props: BasicTableProps) {
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>No</TableCell>
            <TableCell align="left">ShortURL</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {props.allURLS.map((row, index) => (
            <TableRow
              key={index}
              sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                {index+1}
              </TableCell>
              <TableCell align="left"><a href={row.ShortURL}>{row.ShortURL}</a></TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}