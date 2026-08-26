export namespace main {

	export class GetLinkResponse {
	    manifestLink: string;
	    downloadLinks: string[];

	    static createFrom(source: any = {}) {
	        return new GetLinkResponse(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manifestLink = source["manifestLink"];
	        this.downloadLinks = source["downloadLinks"];
	    }
	}

}
