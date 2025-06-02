export class TrackData {
    artists: string[]
    album: string
    track: string
    listenedAt: number
    listenCount: number
    listenTime: number

    constructor(artists: string[], album: string, track: string, listenedAt: number, listenCount: number, listenTime: number) {
        this.artists = artist;
        this.album = album;
        this.track = track;
        this.listenedAt = listenedAt;
        this.listenCount = listenCount;
        this.listenTime = listenTime;
    }
}