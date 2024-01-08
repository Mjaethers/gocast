import { Delete, getData, postData, putData, Time } from "../global";
import { StreamableMapProvider } from "./provider";
import {Progress} from "../api/progress";

export class StreamPlaylistProvider extends StreamableMapProvider<number, StreamPlaylistEntry[]> {
    protected async fetcher(streamId: number): Promise<StreamPlaylistEntry[]> {
        const result = await StreamPlaylist.get(streamId);
        return result
            .map((e) => {
                e.startDate = new Date(e.start);
                return e;
            })
            .sort((a, b) => (a.startDate < b.startDate ? -1 : 1));
    }
}

export interface StreamPlaylistEntry {
    streamId: number;
    courseSlug: string;
    streamName: string;
    liveNow: boolean;
    Progress: Progress; // todo call to progress.ts constructor
    start: string;
    createdAt: string;

    // Client Generated
    startDate: Date;
};

const StreamPlaylist = {
    get: async function (streamId: number): Promise<StreamPlaylistEntry[]> {
        const resp = await getData(`/api/stream/${streamId}/playlist`);
        if (!resp.ok) {
            throw Error(resp.statusText);
        }
        return resp.json();
    },
};
