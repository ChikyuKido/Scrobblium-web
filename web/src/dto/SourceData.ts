export class SourceData {
    name: string
    provider: number
    apiKey: string
    scrobbles: number
    lastScrobble: number

    constructor(name: string, provider: number, apiKey: string, scrobbles: number, lastScrobble: number) {
        this.name = name;
        this.provider = provider;
        this.apiKey = apiKey;
        this.scrobbles = scrobbles;
        this.lastScrobble = lastScrobble;
    }
}