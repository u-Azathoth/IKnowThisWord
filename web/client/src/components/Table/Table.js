import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import MaterialTable from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import MaterialTableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import TableRow from "./templates/table-row";

const useStyles = makeStyles({
    root: {
        width: '100%',
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
});

export default function Table(props) {
    const classes = useStyles();
    const { headers, rows = [], keyField, actions = [] } = props;

    return (
        <Paper className={classes.root}>
            <MaterialTable className={classes.table} aria-label="simple table">
                <TableHead>
                    <MaterialTableRow>
                        {headers.map(header => <TableCell align="justify" key={header}>{header}</TableCell>)}
                    </MaterialTableRow>
                </TableHead>

                <TableBody>
                    {rows.length
                        ? (rows.map(row => (
                            <MaterialTableRow key={row[keyField]}>
                                <TableRow row={row} actions={actions}/>
                            </MaterialTableRow>
                        )))
                        : (
                            <MaterialTableRow>
                                <TableCell colSpan={12} align="center" size="medium">Nothing to show...</TableCell>
                            </MaterialTableRow>
                        )
                    }
                </TableBody>
            </MaterialTable>
        </Paper>
    );
}
