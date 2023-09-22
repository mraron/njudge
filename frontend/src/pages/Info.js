import ProfileSideBar from '../components/concrete/other/ProfileSidebar'
import {SVGCode, SVGCopy, SVGCorrectSimple} from "../svg/SVGs";
import SVGTitleComponent from "../svg/SVGTitleComponent";
import RoundedTable from "../components/container/RoundedTable";
import React, {useState} from "react";
import {useTranslation} from "react-i18next";
import CopyableCode from "../components/util/copy/CopyableCode";
import CopyableCommand from "../components/util/copy/CopyableCommand";

function CompilerOption({lang, command}) {
    const handleCopy = () => {
        navigator.clipboard.writeText(command)
        window.flash("info.successful_copy", "success")
    }
    return (
        <tr className={`divide-x divide-default`}>
            <td className="padding-td-default whitespace-nowrap">
                {lang}
            </td>
            <td className="padding-td-default text-white">
                <CopyableCommand text={command}/>
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