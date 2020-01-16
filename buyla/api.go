package Buyla

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Api struct {
	data *MetaData
}

func NewApi(d *MetaData) *Api {
	return &Api{data: d}
}

func getOrderParameter(r *http.Request) (gid, uid string) {
	pathParams := mux.Vars(r)
	if val, ok := pathParams["gid"]; ok {
		gid = string(val)
	}
	if val, ok := pathParams["uid"]; ok {
		uid = string(val)
	}
	return
}

func (this *Api) checkGroupIsOpen(gid string) error {
	if _, ok := this.data.Groups[gid]; !ok {
		return fmt.Errorf("Is not open")
	}
	if !this.data.Groups[gid].IsOpen {
		return fmt.Errorf("Is not open")
	}
	return nil
}

func (this *Api) HandleGetOrders(w http.ResponseWriter, r *http.Request) {
	gid, _ := getOrderParameter(r)
	err := this.checkGroupIsOpen(gid)
	if err != nil {
		//TODO return error message
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	records := this.data.Groups[gid].Records
	json.NewEncoder(w).Encode(records)
	return
}

func (this *Api) checkUid(gid, uid string) error {
	return nil
}

func (this *Api) HandleGetCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	gid, uid := getOrderParameter(r)

	err := this.checkGroupIsOpen(gid)
	if err != nil {
		//TODO return error message
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}

	err = this.checkUid(gid, uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	// for debug
	//json.NewEncoder(w).Encode(&isValid)

	return
}

func (this *Api) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	gid, uid := getOrderParameter(r)
	group := this.data.Groups[gid]

	err := this.checkGroupIsOpen(gid)
	if err != nil {
		//TODO return error message
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}

	err = this.checkUid(gid, uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group.Records[uid].Goods)
	return
}

func (this *Api) HandlePutOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{")
	fmt.Fprintf(w, `"method":"%s"`, r.Method)
	fmt.Fprintf(w, "}")
	return
}

func (this *Api) HandlePostOrder(w http.ResponseWriter, r *http.Request) {
	gid, uid := getOrderParameter(r)
	err := this.checkGroupIsOpen(gid)
	if err != nil {
		//TODO return error message
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}

	err = this.checkUid(gid, uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}

	if r.Body == nil {
		//TODO return error message
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var record Record
	err = json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err)
		return
	}

	group := this.data.Groups[gid]
	if _, ok := group.Records[uid]; !ok {
		group.Records[uid] = NewRecord(record.UserProfile)
	}

	for _, v := range record.Goods {
		group.Records[uid].Goods = append(group.Records[uid].Goods, v)
		log.Println("[ADD]", gid, record.UserProfile.DisplayName, v)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
	return
}

func (this *Api) HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	gid, uid := getOrderParameter(r)
	err := this.checkGroupIsOpen(gid)
	if err != nil {
		//TODO return error message
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}

	err = this.checkUid(gid, uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}

	group := this.data.Groups[gid]
	if _, ok := group.Records[uid]; !ok {
		w.WriteHeader(http.StatusOK)
		return
	}
	log.Println("[DELETE]", gid, group.Records[uid].UserProfile.DisplayName)
	group.Records[uid] = nil
	delete(group.Records, uid)
	w.WriteHeader(http.StatusOK)
	return
}

func (this *Api) HandlePatchOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{")
	fmt.Fprintf(w, `"method":"%s"`, r.Method)
	fmt.Fprintf(w, "}")
	return
}
