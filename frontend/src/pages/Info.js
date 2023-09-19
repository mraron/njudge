import ProfileSideBar from '../components/concrete/other/ProfileSidebar'
import {SVGCode, SVGCopy, SVGCorrectSimple} from "../svg/SVGs";
import SVGTitleComponent from "../svg/SVGTitleComponent";
import RoundedTable from "../components/container/RoundedTable";
import React, {useState} from "react";
import {useTranslation} from "react-i18next";

function CompilerOption({lang, command}) {
    const {t} = useTranslation()
    const handleCopy = () => {
        navigator.clipboard.writeText(command)
        window.flash(t("info.successful_copy"), "success")
    }
    return (
        <tr className={`divide-x divide-default `}>
            <td className="padding-td-default whitespace-nowrap">
                {lang}
            </td>
            <td className="padding-td-default text-white">
                <div className="flex items-center">
                    <button
                        className="h-9 w-9 mr-2 rounded-md border-1 bg-grey-800 border-grey-725 hover:bg-grey-775 transition duration-200 relative"
                        onClick={handleCopy}>
                        <SVGCopy
                            cls={`absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-4 h-4 `}/>
                    </button>
                    <div className="flex items-center px-3 py-2 border-1 border-grey-725 rounded-md bg-grey-875">
                        <pre>{command}</pre>
                    </div>
                </div>
            </td>
        </tr>
    );
}

function InfoTable() {
    const {t} = useTranslation()
    const compilerOptions = [
        ["C++ (11 / 14 / 17)", "g++ -std=c++<verzió> -O2 -static -DONLINE_JUDGE main.cpp"],
        ["C#", "/usr/bin/mcs -out:main.exe -optimize+ main.cs"],
        ["Go", "/usr/bin/gccgo main.go"],
        ["Java", "/usr/bin/javac main.java"],
        ["Pascal", "/usr/bin/fpc -Mobjfpc -O2 -Xss main.pas"],
        ["PyPy3", "/usr/bin/pypy3 main.py"],
        ["Python3", "/usr/bin/python3 main.py"]
    ];
    const compilerOptionElems = compilerOptions.map((item, index) =>
        <CompilerOption lang={item[0]} command={item[1]} key={index}/>
    );
    const titleComponent = <SVGTitleComponent title={t("info.compiler_options")} svg={<SVGCode cls="w-7 h-7 mr-2"/>}/>
    return (
        <RoundedTable titleComponent={titleComponent}>
            <tbody className="divide-y divide-default text-sm">
            {compilerOptionElems}
            </tbody>
        </RoundedTable>
    );
}

function Info({ data }) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar/>
                </div>
                <div className="w-full px-4 lg:pl-3 overflow-x-auto">
                    <InfoTable/>
                </div>
            </div>
        </div>
    );
}

export default Info;