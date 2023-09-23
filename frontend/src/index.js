import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import UserProvider from "./contexts/user/UserProvider";
import JudgeDataProvider from "./contexts/judgeData/JudgeDataProvider";
import './index.css';
import "./i18n";

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
    <UserProvider>
        <JudgeDataProvider>
            <App/>
        </JudgeDataProvider>
    </UserProvider>
);