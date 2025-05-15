export class TrackData {
    artist: string
    album: string
    track: string
    listenedAt: number
    listenCount: number
    listenTime: number

    constructor(artist: string, album: string, track: string, listenedAt: number, listenCount: number, listenTime: number) {
        this.artist = artist;
        this.album = album;
        this.track = track;
        this.listenedAt = listenedAt;
        this.listenCount = listenCount;
        this.listenTime = listenTime;
    }
}