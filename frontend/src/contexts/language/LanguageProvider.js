import LanguageContext from "./LanguageContext";
import {useState} from "react";
import Cookies from "js-cookie";

function LanguageProvider({children}) {
    const [language, setLanguage] = useState(Cookies.get("language") || "hu")
    const storeLanguage = (lang) => {
        setLanguage(lang)
        Cookies.set("language", lang, {secure: true})
    }
    return (
        <LanguageContext.Provider value={{language, storeLanguage}}>
            {children}
        </LanguageContext.Provider>
    )
}

export default LanguageProvider