package main

func lookupMember(member iRaiserMember) (int, error) {
	if resp, err := SearchContact(member.Mail) ; err != nil {
		return 0, err
	} else {
		return resp.Count, nil
	}
}
