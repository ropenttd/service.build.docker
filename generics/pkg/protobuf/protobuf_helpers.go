package protobuf

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

// pbUnmarshalRequest fills the given Protobuf message with the data from the request.
// Tolerant of JSON requests and Protobuf requests.
func PbUnmarshalHttpRequest(pb proto.Message, r *http.Request) (err error) {
	if r.Header.Get("Content-Type") == "application/json" {
		// Attempt to unmarshal as a JSON request.
		err = jsonpb.Unmarshal(r.Body, pb)
	} else {
		// Otherwise, attempt unmarshalling as Protobuf data.
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		err = proto.Unmarshal(data, pb)
	}
	if err != nil {
		return err
	}
	return nil
}

// Shorthand for pbMarshalHttpRepsonseWithStatus, except returns 200.
func PbMarshalHttpResponse(pb *proto.Message, r *http.Request, w http.ResponseWriter) (err error) {
	return PbMarshalHttpResponseWithStatus(pb, r, w, http.StatusOK)
}

// pbMarshalHttpResponseWithStatus pushes the given Protobuf message back to the client, based on variables set in the request.
// Tolerant of JSON requests and Protobuf requests.
func PbMarshalHttpResponseWithStatus(pb *proto.Message, r *http.Request, w http.ResponseWriter, s int32) (err error) {

	var responseBody []byte

	if r.Header.Get("Content-Type") == "application/json" {
		// Attempt to marshal as JSON.
		marshaler := jsonpb.Marshaler{
			EnumsAsInts:  true,
			EmitDefaults: false,
			OrigName:     false,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(int(s))
		// The jsonpb marshaler emits straight to an io.writer.
		err := marshaler.Marshal(w, *pb)
		if err != nil {
			// This is arguably fatal?
			log.Printf("⁉️ Failed to marshal Protobuf response: %s", err)
			return err
		}
	} else {
		// Otherwise, attempt unmarshalling as Protobuf data.

		// This has to be unmarshaled into a []byte array because that's how Protobuf rolls.
		responseBody, err = proto.Marshal(*pb)
		// Should this status be overridable?
		w.WriteHeader(int(s))
		_, err = w.Write(responseBody)
		if err != nil {
			log.Printf("⁉️ Failed to marshal Protobuf response: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	}

	// We can't have a single error catch here because the two methods use different ways of writing to the response.
	// Nor can we do any logic. Instead, it's all handled in their respective if loops.

	return nil

}
