import React from 'react';
import './App.css';
import { BrowserRouter, Redirect, Route, Switch } from "react-router-dom";
import Vocabulary from "../../pages/Vocabulary/Vocabulary";
import Revision from "../../pages/Revision/Revision";
import Navbar from "../Navbar/Navbar";

const routes = [
    { path: '/vocabulary', component: Vocabulary },
    { path: '/revision', component: Revision },
];

function App() {
    return (
        <BrowserRouter>
            <Navbar/>

            <Switch>
                <Route path="/" exact>
                    <Redirect to="/vocabulary"/>
                </Route>

                {routes.map(({path, component}) => (
                    <Route key={path} path={path} component={component}/>
                ))}
            </Switch>
        </BrowserRouter>
    );
}

export default App;
