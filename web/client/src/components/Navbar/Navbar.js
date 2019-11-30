import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import { useHistory, useLocation } from "react-router-dom";

function a11yProps(index) {
    return {
        id: `nav-tab-${index}`,
        'aria-controls': `nav-tabpanel-${index}`,
    };
}

function LinkTab(props) {
    return (
        <Tab
            component="a"
            onClick={event => {
                event.preventDefault();
            }}
            {...props}
        />
    );
}

const useStyles = makeStyles(() => ({
    root: {
        flexGrow: 1,
    },
}));

export default function Navbar() {
    const classes = useStyles();
    const history = useHistory();
    const location = useLocation();

    const [value, setValue] = React.useState(location.pathname);

    const handleChange = (event, newValue) => {
        setValue(newValue);
        history.push(newValue);
    };

    return (
        <div className={classes.root}>
            <AppBar position="static">
                <Tabs
                    variant="fullWidth"
                    value={value}
                    onChange={handleChange}
                >
                    <LinkTab label="Vocabulary" value="/vocabulary" {...a11yProps(0)} />
                    <LinkTab label="Revision" value="/revision" {...a11yProps(1)} />
                </Tabs>
            </AppBar>
        </div>
    );
}
