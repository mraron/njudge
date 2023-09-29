import { apiRoute } from "../config/RouteConfig";
import fetchWithCredentials from "./fetchWithCredentials";

async function submitSolution({ problem, language, file, submissionCode }) {
    const formData = new FormData();
    formData.append("problem", problem);
    formData.append("language", language);
    if (file) {
        formData.append("file", file);
    }
    if (submissionCode) {
        formData.append("submissionCode", submissionCode);
    }
    const requestOptions = {
        method: "POST",
        body: formData,
    };
    const response = await fetchWithCredentials(
        apiRoute("/problemset/main/submit/"),
        requestOptions,
    );
    return response.ok;
}

export default submitSolution;
