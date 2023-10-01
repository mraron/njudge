import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import ProfileSideBar from "../components/concrete/other/ProfileSidebar";
import SVGTitleComponent from "../components/svg/SVGTitleComponent";
import RoundedTable from "../components/container/RoundedTable";
import CopyableCommand from "../components/util/copy/CopyableCommand";

function CompilerOption({ lang, command }) {
    return (
        <tr className={`divide-x divide-dividecol`}>
            <td className="padding-td-default whitespace-nowrap w-32 text-center">
                {lang}
            </td>
            <td style={{ maxWidth: 0 }}>
                <CopyableCommand text={command} cls="border-0 rounded-none" />
            </td>
        </tr>
    );
}

function InfoTable() {
    const { t } = useTranslation();
    const compilerOptions = [
        ["C++", "g++ -std=c++<verzió> -O2 -static -DONLINE_JUDGE main.cpp"],
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
    return (
        <RoundedTable title={t("info.compiler_options")}>
            <thead className="bg-framebgcol">
                <tr className="divide-x divide-bordercol">
                    <th className="padding-td-default">Nyelv</th>
                    <th className="padding-td-default">Parancs</th>
                </tr>
            </thead>
            <tbody className="divide-y divide-dividecol text-sm">
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
