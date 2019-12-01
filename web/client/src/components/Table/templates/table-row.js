import TableCell from "@material-ui/core/TableCell";
import React from "react";

export default function TableRow(props) {
    const { row = {}, actions = [] } = props;
    let actionRow = null;

    const keysRows = Object.entries(row).map(([ key, value ]) =>
        <TableCell align="justify" key={key}>{value}</TableCell>
    );

    if (actions.length) {
        const [ { title, fn } ] = actions;
        actionRow = (
            <TableCell align="justify">
                <button onClick={fn.bind(null, row)}>{title}</button>
            </TableCell>
        );
    }

    return (
        <React.Fragment>
            {keysRows}
            {(actionRow ? actionRow : ''.trim())}
        </React.Fragment>
    )
}
