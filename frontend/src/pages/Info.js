import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCode } from "@fortawesome/free-solid-svg-icons";
import ProfileSideBar from "../components/concrete/other/ProfileSidebar";
import SVGTitleComponent from "../components/svg/SVGTitleComponent";
import RoundedTable from "../components/container/RoundedTable";
import CopyableCommand from "../components/util/copy/CopyableCommand";

function CompilerOption({ lang, command }) {
    return (
        <tr className={`divide-x divide-default`}>
            <td className="padding-td-default whitespace-nowrap">{lang}</td>
            <td className="padding-td-default">
                <CopyableCommand text={command} />
            </td>
        </tr>
    );
}

function InfoTable() {
    const { t } = useTranslation();
    const compilerOptions = [
        [
            "C++ (11 / 14 / 17)",
            "g++ -std=c++<verzió> -O2 -static -DONLINE_JUDGE main.cpp",
        ],
        ["C#", "/usr/bin/mcs -out:main.exe -optimize+ main.cs"],
        ["Go", "/usr/bin/gccgo main.go"],
        ["Java", "/usr/bin/javac main.java"],
        ["Pascal", "/usr/bin/fpc -Mobjfpc -O2 -Xss main.pas"],
        ["PyPy3", "/usr/bin/pypy3 main.py"],
        ["Python3", "/usr/bin/python3 main.py"],
    ];
    const compilerOptionElems = compilerOptions.map((item, index) => (
        <CompilerOption lang={item[0]} command={item[1]} key={index} />
    ));
    const titleComponent = (
        <SVGTitleComponent
            title={t("info.compiler_options")}
            svg={<FontAwesomeIcon icon={faCode} className="w-5 h-5 mr-2" />}
        />
    );
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
                    <ProfileSideBar />
                </div>
                <div className="w-full px-4 lg:pl-3 overflow-x-auto">
                    <InfoTable />
                </div>
            </div>
        </div>
    );
}

export default Info;
