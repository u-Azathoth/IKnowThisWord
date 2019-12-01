import React, { useEffect, useState } from "react";
import config from "../../config/environments.dev";
import Table from "../../components/Table/Table";

export default function Vocabulary() {
    const [ words, setWords ] = useState([]);

    useEffect(() => {
        (async () => setWords(
            await fetch(`${config.api.URL}`).then(response => response.json())
        ))();
    }, []);

    const headers = [ 'Id', 'Word', 'Meaning', 'Recognition Rate', '' ];
    const actions = [
        { title: 'Delete', fn: (row) => console.log('Delete row', row) },
    ];

    return <Table headers={headers} rows={words} keyField="id" actions={actions}/>
}
