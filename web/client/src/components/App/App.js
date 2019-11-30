import React from 'react';
import './App.css';
import { BrowserRouter, Redirect, Route, Switch } from "react-router-dom";
import Vocabulary from "../../pages/Vocabulary/Vocabulary";
import Revision from "../../pages/Revision/Revision";
import Navbar from "../Navbar/Navbar";

function App() {
    return (
        <BrowserRouter>
            <Navbar/>

            <Switch>
                <Route path="/" exact>
                    <Redirect to="/vocabulary"/>
                </Route>

                <Route path="/vocabulary" component={Vocabulary} exact/>
                <Route path="/revision" component={Revision} exact/>
            </Switch>
        </BrowserRouter>
    );
}

export default App;
