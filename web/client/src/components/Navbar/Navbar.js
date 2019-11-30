import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import { useHistory, useLocation } from "react-router-dom";
import { makeStyles, withStyles } from '@material-ui/core/styles';

const a11yProps = index => ({
    id: `nav-tab-${index}`,
    'aria-controls': `nav-tabpanel-${index}`,
});

const StyledTabs = withStyles({
    indicator: {
        backgroundColor: 'green'
    },
})(props => <Tabs {...props} TabIndicatorProps={{children: <div/>}}/>);

const StyledTab = withStyles(theme => ({
    root: {
        '&$selected': {
            color: '#fff',
            fontWeight: theme.typography.fontWeightMedium,
        },
    },
    selected: {},
}))(props => <Tab disableRipple {...props} />);

const useStyles = makeStyles(() => ({
    root: {
        flexGrow: 1,
    },
    wrapper: {
        color: "hsla(0,0%,100%,.7)",
        backgroundColor: "#24292e"
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
                <StyledTabs variant="fullWidth"
                            value={value}
                            onChange={handleChange}
                            className={classes.wrapper}
                >
                    <StyledTab label="Vocabulary" value="/vocabulary" {...a11yProps(0)} />
                    <StyledTab label="Revision" value="/revision" {...a11yProps(1)} />
                </StyledTabs>
            </AppBar>
        </div>
    );
}
