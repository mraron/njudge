import React from "react";
import ReactDOM from "react-dom/client";
import { fas } from "@fortawesome/free-solid-svg-icons";
import { far } from "@fortawesome/free-regular-svg-icons";
import { fab } from "@fortawesome/free-brands-svg-icons";
import { library } from "@fortawesome/fontawesome-svg-core";

import App from "./App";
import UserProvider from "./contexts/user/UserProvider";
import JudgeDataProvider from "./contexts/judgeData/JudgeDataProvider";
import ThemeProvider from "./contexts/theme/ThemeProvider";

import "tw-elements-react/dist/css/tw-elements-react.min.css";
import "./index.css";
import "./i18n";

library.add(fas, far, fab)

const root = ReactDOM.createRoot(document.getElementById("root"))
root.render(
    <UserProvider>
        <JudgeDataProvider>
            <ThemeProvider>
                <App />
            </ThemeProvider>
        </JudgeDataProvider>
    </UserProvider>,
)
