import { useTranslation } from "react-i18next"
import Ranklist from "../../components/concrete/other/Ranklist"
import Pagination from "../../components/util/Pagination"
import RanklistNotPublic from "../../components/concrete/other/RanklistNotPublic"

function ProblemRanklist({ data }) {
    const { t } = useTranslation()
    return (
        <>
            {data.isPublic && (
                <div className="space-y-2">
                    <Ranklist
                        ranklist={data.ranklist}
                        title={t("problem_ranklist.ranklist")}
                    />
                    <Pagination paginationData={data.paginationData} />
                </div>
            )}
            {!data.isPublic && <RanklistNotPublic />}
        </>
    )
}

export default ProblemRanklist
